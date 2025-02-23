package reveal_url

import "net/http"
import net_url "net/url"

const MAX_REQUEST_ATTEMPT = 10

func getResponseRedirectUrl(response *http.Response) *net_url.URL {
  redirectUrl, _ := net_url.Parse(response.Header.Get("Location"))

  if redirectUrl.Host != "" {
    return redirectUrl
  }

  var requestUrl *net_url.URL = response.Request.URL

  redirectUrl.Host = requestUrl.Host
  redirectUrl.Scheme = requestUrl.Scheme

  return redirectUrl
}

func isResponseRedirect(response *http.Response) bool {
  return response.StatusCode >= 300 &&
      response.StatusCode < 400 &&
      len(response.Header.Get("Location")) > 0
}

func goTo(url string) (response *http.Response, err error) {
  client := &http.Client{
    CheckRedirect: func (req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }

  request, err := http.NewRequest("HEAD", url, nil)
  if err != nil {
    return nil, err
  }

  // Some Server seems to reject request missing the User-Agent header
  request.Header.Add("User-Agent", "")

  return client.Do(request)
}

func Reveal(url string, urls *[]string) (*[]string, error) {
  response, err := goTo(url)

  if err != nil {
    return nil, err;
  }

  *urls = append(*urls, url)

  // should Redirect?
  if len(*urls) + 1 > MAX_REQUEST_ATTEMPT || len(response.Header.Get("Location")) == 0 || response.Header.Get("Location") == url {
    return urls, nil;
  }

  isRedirect := isResponseRedirect(response)

  if isRedirect {
    return Reveal(getResponseRedirectUrl(response).String(), urls)
  }

  if len(*urls) > 0 {
    return urls, nil;
  }

  return nil, nil;
}

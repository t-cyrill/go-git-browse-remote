package main

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "strings"

  "github.com/urfave/cli"
)

func git_current_branch() string {
  out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error: git config --get remote.origin.url")
  }
  return strings.TrimSpace(string(out))
}

func git_web_browse(url string) {
  _, err := exec.Command("git", "web--browse", url).Output()
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error: call git web--browse")
  }
  return
}

func git_get_remote_url() string {
  out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error: git config --get remote.origin.url")
  }
  return strings.TrimSpace(string(out))
}

func main() {
  var directory string

  app := cli.NewApp()
  app.Name = "git-browse-remote"
  app.Usage = "open git remote repository on your web browser"
  app.Version = "0.0.1"
  app.Flags = []cli.Flag{
    cli.BoolFlag{
      Name:  "stdout",
      Usage: "prints URL instead of opening browser",
    },
    cli.BoolFlag{
      Name:  "pullrequest, pr",
      Usage: "open pull request URL instead of top",
    },
    cli.StringFlag{
      Name:        "directory",
      Usage:       "change working directory",
      Destination: &directory,
    },
  }
  app.Action = func(c *cli.Context) error {
    var url string
    url_placeholder := "https://github.com/%s"
    directory := c.GlobalString("directory")
    if len(directory) > 0 {
      os.Chdir(directory)
    }

    if c.GlobalBool("pullrequest") {
      url_placeholder = fmt.Sprintf("https://github.com/%%s/pull/%s", git_current_branch())
    }

    is_stdout := c.GlobalBool("stdout")

    patterns := []string{"git@github\\.com:(.*)\\.git", "https?://github.com/(.*)(\\.git)?"}
    remote_url := git_get_remote_url()
    for _, s := range patterns {
      regex_pattern := regexp.MustCompile(s)
      m := regex_pattern.FindStringSubmatch(remote_url)
      if len(m) > 0 {
        url = fmt.Sprintf(url_placeholder, m[1])
        break
      }
    }

    if len(url) == 0 {
      fmt.Println("cannot detect remote url, are you in git repository?")
      return nil
    }

    if is_stdout {
      fmt.Printf("%s\n", url)
    } else {
      git_web_browse(url)
    }
    return nil
  }

  app.Run(os.Args)
}

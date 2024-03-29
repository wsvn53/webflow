# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://github.com/wsvn53/webflow/releases/download/v0.5.0/webflow-darwin-amd64"
  version "v0.5.0"
  sha256 "34264b838c7cfcd87f945d466eb608171f59f43ca7ac27a78f5f847f08d34f79"

  def install
    system "find", ".", "-name", "*webflow*", "-exec", "mv", "{}", "webflow", ";"
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

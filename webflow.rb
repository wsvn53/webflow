# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://github.com/wsvn53/webflow/releases/download/v0.5.0/webflow-darwin-amd64"
  version"v0.5.0"
  sha256"d036a6c42711a667ec5d383083a1ff02c6d07ac91e7bb132d660411518779f49"

  def install
    system "find", ".", "-name", "*webflow*", "-exec", "mv", "{}", "webflow", ";"
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

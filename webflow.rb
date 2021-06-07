# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://git.wsen.me/attachments/ca09eae3-8718-4d9f-a36b-b3698abbc257"
  version "v0.4.1"
  sha256 "6a13d2e058b28eb1b687ee261987845cfc88da3cc1f1cdd74401a47f4256efac"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url"https://git.wsen.me/attachments/92c91255-b6da-4d34-b752-68654b19eaec"
  version"v0.4.7"
  sha256"2c3c9cf2bdfafeeb5afc6f9e08993030738cffa175160fc51a34a654fe1704eb"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

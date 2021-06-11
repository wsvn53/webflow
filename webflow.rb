# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://git.wsen.me/attachments/922490cb-5fde-4c4a-928e-e6520cf2f0cd"
  version "v0.4.5"
  sha256 "3b9f96a01c5c48742df3574c9f6e09ae70c53d46177863853e30f2380a6af3f5"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

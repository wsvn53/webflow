# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://git.wsen.me/attachments/145e6ef3-6edf-4284-a58c-95f28a42b080"
  version "v0.4.3"
  sha256 "ad308b2c1a047bf8e821b54e53cd4a13ffec77e25b91a2c99ebfa3e9940bfc0b"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

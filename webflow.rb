# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://git.wsen.me/attachments/7770809e-2711-4467-bc69-32f631dbd2dd"
  version "v0.4.2"
  sha256 "d40a58e42e7fa6e1910329b726902802da91e57fadc38cc1fe6477884d2a0a50"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

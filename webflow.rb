# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://git.wsen.me/attachments/94fb0123-abcd-45c5-8747-3aa6756c7b64"
  version "v0.4.6"
  sha256 "6562db16e2ff2c152cbb72fabeb3c4e8a314cc71b04ba4592bb603eefcdd4092"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Webflow: web automation tool based on chromedp"
  homepage "https://git.wsen.me/utils/webflow"
  url "https://git.wsen.me/attachments/f6cd854f-44eb-4e2a-9ae5-267364d80a76"
  version "v0.4.4"
  sha256 "f6a3ae950ff19de6f110444124368399c4064301fbd4c2b144a4e6a95c9e9b70"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

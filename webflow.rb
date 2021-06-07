# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
class Webflow < Formula
  desc "Hotcode is an iOS app code hot-reload tool."
  homepage "https://git.wsen.me/utils/webflow"
  url "https://git.wsen.me/attachments/f74423a6-8d75-49a9-8de5-5330414c4236"
  version "v0.4"
  sha256 "35620db50a54b7d92e40359b63f65ae96409c4f697484e607af897ba88f09ec7"

  def install
    system "chmod", "+x", "webflow"
    bin.install "webflow"
  end

  test do
    system "false"
  end
end

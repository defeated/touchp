class Touchp < Formula
  desc      "basically `mkdir -p` + `touch`, written in Go"
  homepage  "https://github.com/defeated/touchp/"
  url       "https://github.com/defeated/touchp/releases/download/0.0.1/touchp-0.1.1-darwin-amd64.zip"
  sha256    "f0275b34d87c624541dd04824815290666bc6de5055632d32d4d154507e27ed0"

  def install
    bin.install "touchp"
  end
end

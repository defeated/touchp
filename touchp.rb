class Touchp < Formula
  url       "https://github.com/defeated/touchp/releases/download/0.0.2/touchp-0.0.2-darwin-amd64.zip"
  homepage  "https://github.com/defeated/touchp/"
  desc      "basically `mkdir -p` + `touch`, written in Go"
  sha256    "a8f4c350d516dfa02840e57f5652ba6a6c51c58676d0852dbf3196275319a962"

  def install
    bin.install "touchp"
  end

  test do
    system "touchp"
  end
end

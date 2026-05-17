class Dockyard < Formula
  desc "TUI for browsing and applying Starship prompt presets"
  homepage "https://github.com/MerrickWykman/dockyard"
  version "1.0.0"

  # After each release, update the version above and replace the SHA256 values
  # below with the checksums printed in the "Print SHA256 checksums" step of
  # the release workflow run.

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/MerrickWykman/dockyard/releases/download/v#{version}/dockyard-mac-arm64"
      sha256 "REPLACE_WITH_SHA256_OF_dockyard-mac-arm64"
    else
      url "https://github.com/MerrickWykman/dockyard/releases/download/v#{version}/dockyard-mac-amd64"
      sha256 "REPLACE_WITH_SHA256_OF_dockyard-mac-amd64"
    end
  end

  on_linux do
    url "https://github.com/MerrickWykman/dockyard/releases/download/v#{version}/dockyard-linux-amd64"
    sha256 "REPLACE_WITH_SHA256_OF_dockyard-linux-amd64"
  end

  def install
    if OS.mac?
      bin.install "dockyard-mac-arm64" => "dockyard" if Hardware::CPU.arm?
      bin.install "dockyard-mac-amd64" => "dockyard" if Hardware::CPU.intel?
    else
      bin.install "dockyard-linux-amd64" => "dockyard"
    end
  end

  test do
    system "#{bin}/dockyard", "--version"
  end
end

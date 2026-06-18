class MacExporter < Formula
  desc "A Prometheus exporter for macOS system metrics using powermetrics and vm_stat commands."
  homepage "https://github.com/rpoetrap/mac-exporter"
  url "https://github.com/rpoetrap/mac-exporter/releases/download/v0.0.0/mac-exporter_Darwin_arm64.tar.gz"
  sha256 "0000000000000000000000000000000000000000000000000000000000000000"
  license "MIT"

  service do
    run opt_bin/"mac-exporter"
    keep_alive true
    log_path var/"log/mac-exporter.log"
    error_log_path var/"log/mac-exporter-error.log"
  end

  def install
    bin.install "mac-exporter"
  end
end

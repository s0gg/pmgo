class Pmgo < Formula
  desc "HTTP/HTTPS Availability Checker tool written in Go"
  homepage "https://github.com/s0gg/pmgo"
  head "https://github.com/s0gg/pmgo.git", branch: "main"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args, "pmgo.go"
  end

  test do
    sytem "#{bin}/pmgo"
  end
end

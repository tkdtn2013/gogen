#!/bin/bash

set -e

APP_NAME="gogen"
VERSION="$1"
RELEASE_REPO="tkdtn2013/gogen"
TAP_REPO_PATH="$2" 
FORMULA_DIR="$TAP_REPO_PATH/Formula"

if [ -z "$VERSION" ]; then
  echo "❌ Usage: ./release_gogen.sh v0.1.0"
  exit 1
fi

echo "🚀 Building $APP_NAME binary..."
GOOS=darwin GOARCH=amd64 go build -o $APP_NAME

echo "📦 Packaging $APP_NAME-$VERSION..."
TAR_NAME="$APP_NAME-macos-$VERSION.tar.gz"
tar -czf $TAR_NAME $APP_NAME

echo "🔒 Calculating SHA256..."
SHA=$(shasum -a 256 $TAR_NAME | awk '{print $1}')
echo "✅ SHA256: $SHA"

# remove build file
rm -f gogen
rm -f $TAR_NAME
echo "✅ remove file"

echo "📝 Creating Formula..."
mkdir -p "$FORMULA_DIR"
cat > "$FORMULA_DIR/${APP_NAME}.rb" <<EOF
class Gogen < Formula
  desc "Run restApi framework gin easily ClI gogen."
  homepage "https://github.com/$RELEASE_REPO"
  url "https://github.com/$RELEASE_REPO/releases/download/$VERSION/$TAR_NAME"
  sha256 "$SHA"
  license "MIT"

  def install
    bin.install "$APP_NAME"
  end

  test do
    system "#{bin}/$APP_NAME", "--help"
  end
end
EOF

echo "✅ Formula created at $FORMULA_DIR/${APP_NAME}.rb"




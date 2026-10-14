// Hand-written: the SDK's version, which scripts/patch_generated.py sends in
// the User-Agent of every request. The release workflow checks that the tag
// matches it.

package internal

// Version is the version of this module, without the leading v.
const Version = "0.1.0"

const userAgent = "roicai-go/" + Version

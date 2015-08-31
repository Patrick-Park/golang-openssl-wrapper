package crypto_test

import (
	. "github.com/ScarletTanager/openssl/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crypto", func() {

	Describe("Using FIPS mode", func() {
		Context("Enabling and disabling FIPS mode", func() {
			It("should return the current setting", func() {
				FIPS_mode_set(1)
				Expect(FIPS_mode()).To(Equal(1))
				FIPS_mode_set(0)
				Expect(FIPS_mode()).To(Equal(0))
			})
		})
	})

	Describe("Managing a cipher context", func() {
		Context("Initializing and freeing the context", func() {
			It("should return indicating success", func() {
				ctx := EVP_CIPHER_CTX_new()
				EVP_CIPHER_CTX_init(ctx)
				Expect(EVP_CIPHER_CTX_cleanup(ctx)).To(Equal(1))
			})
		})
	})
})
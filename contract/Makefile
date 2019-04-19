ed25519::
	$(eval TMPDIR=$(shell mktemp -d))
	solc Ed25519.sol --bin --abi --optimize -o $(TMPDIR)
	cp $(TMPDIR)/Ed25519.abi .
	cp $(TMPDIR)/Ed25519.bin .
	abigen --bin=Ed25519.bin --abi=Ed25519.abi --pkg=ed25519 --out=ed25519.go

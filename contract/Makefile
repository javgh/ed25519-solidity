curve25519::
	$(eval TMPDIR=$(shell mktemp -d))
	solc Curve25519.sol --bin --abi --optimize -o $(TMPDIR)
	cp $(TMPDIR)/Curve25519.abi .
	cp $(TMPDIR)/Curve25519.bin .
	abigen --bin=Curve25519.bin --abi=Curve25519.abi --pkg=curve25519 --out=curve25519.go

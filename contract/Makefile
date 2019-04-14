outpost:
	$(eval TMPDIR=$(shell mktemp -d))
	solc Outpost.sol --bin --abi --optimize -o $(TMPDIR)
	cp $(TMPDIR)/Outpost.abi .
	cp $(TMPDIR)/Outpost.bin .
	abigen --bin=Outpost.bin --abi=Outpost.abi --pkg=outpost --out=outpost.go

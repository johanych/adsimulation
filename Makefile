all: ransomware discovery uac_bypass

ransomware:
	go build github.com/johanych/adsimulation/cmd/ransomware

discovery:
	go build github.com/johanych/adsimulation/cmd/discovery

uac_bypass:
	go build github.com/johanych/adsimulation/cmd/uac_bypass

registry_run:
	go build github.com/johanych/adsimulation/cmd/runkeyregistry
	
gorelease:
	goreleaser release --rm-dist --snapshot 


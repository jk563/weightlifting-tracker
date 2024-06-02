plan: package
	cd infrastructure && just plan

deploy: package
	cd infrastructure && just apply && just test_get

package:
	cd application && just package

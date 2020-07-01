BUILD := /app/build
URI := tcp://0.0.0.0:3000
SERVEOPTS := -s -n -l $(URI)

start:
	npx serve $(SERVEOPTS) $(BUILD)

setup:
	npm install
	npm install
	npm install
	npm audit fix

build_pages:
	npm run build

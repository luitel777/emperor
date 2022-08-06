
run:
	./emperor 8080 false
build:
	@mkdir web/pdf && cd web/pdf && curl -LO https://github.com/mozilla/pdf.js/releases/download/v2.15.349/pdfjs-2.15.349-dist.zip && unzip  pdfjs-2.15.349-dist.zip 
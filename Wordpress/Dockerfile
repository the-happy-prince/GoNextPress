FROM wordpress:php7.1-apache
WORKDIR /var/www/html
COPY . .
EXPOSE 80
RUN pear install PHP_CodeSniffer
RUN phpcs --config-set default_standard WordPress
CMD ["apache2-foreground"]

FROM ruby:2.0.0-p648

RUN gem install sinatra -v 1.4.8

COPY ./wb.rb /app/wb.rb
WORKDIR /app
EXPOSE 9009
RUN chmod +x wb.rb
ENTRYPOINT [ "ruby", "wb.rb" ]

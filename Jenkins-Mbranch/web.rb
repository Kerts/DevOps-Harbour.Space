require 'sinatra'
  require 'json'

  set :port, 9090

  get '/' do
    { name: 'Hello',
      description: 'World',
      Url: request.url }.to_json
  end

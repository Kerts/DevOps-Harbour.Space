#!/usr/bin/env ruby

require 'sinatra'
  require 'json'

  set :bind, '0.0.0.0'
  set :port, 9009

  get '/' do
    { name: 'Hello',
      description: 'World',
      Url: request.url }.to_json
  end

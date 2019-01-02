# Channel
Channel is reference

channels block code till it pull values in and off
range channels blocks till channel is closed

- making channel
- putting values on a channel
- taking values off of a channel
- buffered Channels

## What is channel and for

## Channels block
- they are like runners in a relay race
  - they are synchronized
  - they have to pass/receive the value at the same time
    - just like runners in a relay race have to pass / receive the baton to each other at the same time
      - one runner can’t pass the baton at one moment
      - and then, later, have the other runner receive the baton
    - the baton is passed/received by the runners at the same time
- the value is passed/received synchronously; at the same time


## Buffered Channel
William say try to be away from buffered channel as we want to have well built code and well interact between channel code

## Channel as a type [#]()

## Concurrency [#](https://golang.org/doc/effective_go.html#concurrency)

## Fan in & Fan out
07.fan-in is good design pattern
- Todd’s code
  https://play.golang.org/p/_CyyXQBCHe 
- Rob Pike’s code
  https://play.golang.org/p/buy30qw5MM 
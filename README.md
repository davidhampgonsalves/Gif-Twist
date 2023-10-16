# Gif Twist
> Animate a series of images around a central point.

<p style="text-align:center">
  <img src="https://github.com/davidhampgonsalves/gif-twist/raw/master/examples/m&m.gif" width=1200px>
  <img src="https://github.com/davidhampgonsalves/gif-twist/raw/master/examples/dubai.gif" width=1200px>
</p>

# Explination
Each output frame that Gif Twist creates is a combination of all input frames offset by time.

<p style="text-align:center">
  <img src="https://github.com/davidhampgonsalves/gif-twist/raw/master/examples/example.gif" width=1200px>
</p>

# Warnings
This project was just an experiment and acts as a working example of the technique. It is not performant because it uses masking to merge each image slice into the out rather then iterating over the output pixels and simply selecting its value from one of the source images.

## Usage (from a video source)
Install go-lang, [avconv](https://libav.org/avconv.html) & [gifsicle](https://www.lcdf.org/gifsicle/).
```
avconv -i time.mp4 -r 1.5 -an -y 'inputs/%04d.png'
go run giftwist.go
gifsicle --delay=4 --optimize --loop out/*.gif > ../out.gif
```

# Credit
I saw this technique manually generated in this [reddit post](https://www.reddit.com/r/gifs/comments/4xdfa9/timescape_halls_harbour_nova_scotia/) by Spiritgreen.

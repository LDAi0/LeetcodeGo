from turtle import *
screensize(2000,2000)
tracer(0)
m=50

left(90)

for k in range(0,2):
    forward(20*m)
    left(270)
    forward(12*m)
    right(90)
penup()
forward(9*m)
right(90)
left(7)
right(90)
pendown()
for k in range(0,2):
    forward(13*m)
    right(90)
    forward(6*m)
    right(90)
penup()

for x in range(-100,100):
    for y in range(-100,100):
        setpos(x*m,y*m)
        dot(5,'red')
done()
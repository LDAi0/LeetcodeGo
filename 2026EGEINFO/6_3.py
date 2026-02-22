from turtle import *
screensize(2000,2000)
tracer(0)
m=50

left(90)

for i in range(0,2):
    forward(14*m)
    left(270)
    backward(12*m)
    right(90)
penup()
forward(9*m)
right(90)
backward(7*m)
left(90)
pendown()
for i in range(0,2):
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
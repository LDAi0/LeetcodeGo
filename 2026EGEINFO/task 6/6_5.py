from turtle import *
screensize(2000,2000)
tracer(0)
m=50

left(50)

right(90)

for i in range(9):
    forward(5*m)
    right(120)

penup()

for x in range(-100,100):
    for y in range(-100,100):
        setpos(x*m,y*m)
        dot(5,'red')
done()
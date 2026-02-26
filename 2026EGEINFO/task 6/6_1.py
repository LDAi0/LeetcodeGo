from turtle import *
screensize(2000,2000)
tracer(0)
m=50

left(90)

for i in range(4):
    forward(10*m)
    right(90)
right(30)
for i in range(5):
    forward(6*m)
    right(60)
    forward(6*m)
    right(120)
penup()

for x in range(-100,100):
    for y in range(-100,100):
        setpos(x*m,y*m)
        dot(5,'red')
done()
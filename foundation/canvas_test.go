package foundation

import "testing"

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	if c.Width != 10 || c.Height != 20 {
		t.Error("NewCanvas is wrong")
	}

	for _, v := range c.Pixels {
		if v.Red != 0.0 || v.Green != 0.0 || v.Blue != 0.0 {
			t.Error("Pixels are wrong")
		}
	}
}

func TestWriteRead(t *testing.T) {
	c := NewCanvas(10, 20)
	clr := Color { 0.5, 0.5, 0.5 }
	c.Write(3, 3, clr)

	r := c.Read(3, 3)
	if r != clr {
		t.Error("Write/Read mismatch in canvas")
	}
}



type projectile struct {
	position Vector
	velocity Vector
}

type environment struct {
	gravity Vector
	wind Vector
}

func tick(env environment, proj projectile) projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add(env.wind)
	return projectile { position, velocity }
}

func TestProjectile(t *testing.T) {
	start := NewPoint(0, 1, 0)
	velocity := NewVector(1, 1.8, 0).Normalize().Mult(11.25)
	p := projectile { start, velocity }
	gravity := NewVector(0, -0.1, 0)
	wind := NewVector(-0.01, 0, 0)
	e := environment { gravity, wind }
	c := NewCanvas(900, 550)

	clr := Color { 1.0, 0.5, 0.0 }
	for p.position.Y > 0.0 {
		c.Write(int(p.position.X), int(p.position.Y), clr)

		p = tick(e, p)
	}

	c.Save("d:\\toto.png")
}
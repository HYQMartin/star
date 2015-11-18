package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["FirstStep"] = `
			<div align="middle">
				<nav>
  				<ul class="pagination pagination-lg">
   				<li>
      				<a href="#" aria-label="Previous">
        			<span aria-hidden="true">&laquo;</span>
      				</a>
    			</li>
    			<li><a href="/first">第一阶段(1.27-2.4)：</a></li>
    			<li><a href="/second">第二阶段(2.12-~)：</a></li>
    			<li>
      				<a href="#" aria-label="Next">
        			<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>
  				</ul>
				</nav>
			</div>
	`
	c.Data["Author"] = "Mar组"
	c.Data["Time"] = time.Now().Format(time.Stamp)
	c.Data["IsHome"] = true

	namestruct, err := c.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		c.Data["UserName"] = name

	}

	c.Data["IsLogin"] = CheckAccount(c.Ctx)
	c.TplNames = "Home.html"

}

type FirstController struct {
	beego.Controller
}

func (c *FirstController) Get() {

	c.Data["FirstStep"] = `
			<div align="middle">
				<nav>
  				<ul class="pagination pagination-lg">
   				<li>
      				<a href="/" aria-label="Previous">
        			<span aria-hidden="true">&laquo;</span>
      				</a>
    			</li>
    			<li class="active"><a href="/first">第一阶段(1.27-2.4)：</a></li>
    			<li><a href="/second">第二阶段(2.12-~)：</a></li>
    			<li>
      				<a href="/second" aria-label="Next">
        			<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>
  				</ul>
				</nav>
			</div>
				<h4 class="text-muted">起始于 2015-1-27</h4>
				
				<p>2-5：2-6回深圳  暂停</p>
	`
	c.Data["Author"] = "Mars组"
	c.Data["Time"] = time.Now().Format(time.Stamp)
	c.Data["IsHome"] = true
	c.TplNames = "Home.html"

}

type SecondController struct {
	beego.Controller
}

func (c *SecondController) Get() {
	c.Data["FirstStep"] = `
			<div align="middle">
				<nav>
  				<ul class="pagination pagination-lg">
   				<li>
      				<a href="/first" aria-label="Previous">
        			<span aria-hidden="true">&laquo;</span>
      				</a>
    			</li>
    			<li><a href="/first">第一阶段(1.27-2.4)：</a></li>
    			<li class="active"><a href="/second">第二阶段(2.12-~)：</a></li>
    			<li class="disabled">
      				<a href="#" aria-label="Next">
        			<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>
  				</ul>
				</nav>
			</div>
				<h4 class="text-muted">起始于 2015-1-27</h4>
				
				<br></br> <br></br> <br></br><br></br><br></br><br></br><br></br><br></br>
	`
	c.Data["Author"] = "Mars组"
	c.Data["Time"] = time.Now().Format(time.Stamp)
	c.Data["IsHome"] = true
	c.TplNames = "Home.html"

}

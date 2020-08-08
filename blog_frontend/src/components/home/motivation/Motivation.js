import React from "react";
import classes from "./Motivation.module.scss";
import go from "../../../assets/golang.png";
const Motivation = () => {
	return (
		<div className="sub-container" id="motivation">
			<h1 className="text-ac mb-1">
				<span className="text-ul">Motivation</span>
			</h1>
			<div className={classes.motivation}>
				<p className={`paragraph ${classes.motivation_para}`}>
					Learning is a never ending process and in the modern world of
					technology one needs to pick its direction and be very clear about it
					instead of trying to learn everything on the go. As a professional I
					spent several years with PHP with all the stuff (Class, Trait,
					Exception, XDebug) it has to offer still language felt lacking in some
					of the most important areas (concurrency, strict typing) to make a
					more robust web application. JavaScript is everywhere and it's awesome
					for server side application to with its non-blocking model but one
					need to keep in mind that its single threaded and it don't serve
					requests in parallel so node I/O may not block the requests but code
					with sleep to particular request does blocks the server. With all this
					mind i started my journey with golang and so far it's been awesome
					golang offers to do things in one way (goroutines, channels,
					statically typed, Cyclic dependencies, Pass By Value) no need to run
					into Functional vs Class-Components discussion as JavaScript community
					did to JavaScript
				</p>
				<div className={classes.motivation_img}>
					<img src={go} alt="go power" />
				</div>
			</div>
		</div>
	);
};

export default Motivation;

import React from "react";
import classes from "./Button.module.scss";
const Button = (props) => (
	<button
		disabled={props.disabled}
		className={[
			classes.btn,
			classes[props.btnType],
			classes[props.btnSize],
			props.classes,
		].join(" ")}
		type={props.type}
		onClick={props.clicked}
	>
		{props.children}
	</button>
);
export default Button;

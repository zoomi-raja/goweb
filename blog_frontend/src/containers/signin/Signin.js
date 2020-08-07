import React from "react";
import Button from "../../components/ui/button/Button";
const onClick = () => {
	window.alert("yet to implement");
};
const Signin = () => {
	return (
		<Button clicked={onClick} btnType="btn--primary">
			Sign in
		</Button>
	);
};

export default Signin;

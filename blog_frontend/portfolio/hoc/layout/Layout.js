import React, { Fragment } from "react";
import Header from "../../components/header/Header";
import Footer from "../../components/footer/Footer";

const Layout = (props) => {
	return (
		<Fragment>
			<Header {...props} />
			{props.children}
			<Footer />
		</Fragment>
	);
};

export default Layout;

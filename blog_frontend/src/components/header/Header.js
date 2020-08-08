import React, { useEffect, useRef, useState } from "react";
import logo from "../../assets/logo.svg";
import classes from "./Header.module.scss";
import Nav from "./nav/Nav";
import Signin from "../../containers/signin/Signin";
const Header = () => {
	const [isSticky, setSticky] = useState(false);
	const ref = useRef(null);
	const handleScroll = () => {
		if (ref.current) {
			setSticky(ref.current.getBoundingClientRect().top <= 0);
		}
	};
	useEffect(() => {
		window.addEventListener("scroll", handleScroll);

		return () => {
			window.removeEventListener("scroll", () => handleScroll);
		};
	}, []);
	let headerClass = isSticky
		? `${classes.header} ${classes.sticky}`
		: classes.header;
	return (
		<header className={`sub-container ${headerClass}`} ref={ref}>
			<nav className={classes.nav_holder}>
				<div className={classes.logo}>
					<img src={logo} alt="react baby" />
				</div>
				<Nav />
			</nav>
			<Signin />
		</header>
	);
};

export default Header;

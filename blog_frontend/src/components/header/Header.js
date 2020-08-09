import React, { useEffect, useRef, useState } from "react";
import logo from "../../assets/logo.svg";
import classes from "./Header.module.scss";
import Nav from "./nav/Nav";
import Signin from "../../containers/signin/Signin";
import { animateScroll as scroll } from "react-scroll";
const scrollToTop = () => {
	scroll.scrollToTop({ smooth: true, duration: 300 });
};
const Header = () => {
	const [isSticky, setSticky] = useState(false);
	const [isDark, setIsDark] = useState(false);
	const changeTheme = () => {
		if (!isDark) {
			document.body.classList.add("dark");
		} else {
			document.body.classList.remove("dark");
		}
		setIsDark(!isDark);
	};
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
	let showGoUp = isSticky ? `${classes.goUp} ${classes.show}` : classes.goUp;
	return (
		<header className={`sub-container ${headerClass}`} ref={ref}>
			<div className={showGoUp} onClick={scrollToTop}>
				&uarr;
			</div>
			<nav className={classes.nav_holder}>
				<div className={classes.logo}>
					<img src={logo} alt="react baby" onClick={changeTheme} />
				</div>
				<Nav />
			</nav>
			<Signin />
		</header>
	);
};

export default Header;

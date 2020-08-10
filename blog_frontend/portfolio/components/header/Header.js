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
					<svg
						onClick={changeTheme}
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
					>
						<path d="M12 9c1.654 0 3 1.346 3 3s-1.346 3-3 3v-6zm0-2c-2.762 0-5 2.238-5 5s2.238 5 5 5 5-2.238 5-5-2.238-5-5-5zm-4.184-.599l-3.593-3.594-1.415 1.414 3.595 3.595c.401-.537.876-1.013 1.413-1.415zm4.184-1.401c.34 0 .672.033 1 .08v-5.08h-2v5.08c.328-.047.66-.08 1-.08zm5.598 2.815l3.595-3.595-1.414-1.414-3.595 3.595c.537.402 1.012.878 1.414 1.414zm-12.598 4.185c0-.34.033-.672.08-1h-5.08v2h5.08c-.047-.328-.08-.66-.08-1zm11.185 5.598l3.594 3.593 1.415-1.414-3.594-3.593c-.403.536-.879 1.012-1.415 1.414zm-9.784-1.414l-3.593 3.593 1.414 1.414 3.593-3.593c-.536-.402-1.011-.877-1.414-1.414zm12.519-5.184c.047.328.08.66.08 1s-.033.672-.08 1h5.08v-2h-5.08zm-6.92 8c-.34 0-.672-.033-1-.08v5.08h2v-5.08c-.328.047-.66.08-1 .08z" />
					</svg>
				</div>
				<Nav />
			</nav>
			<Signin />
		</header>
	);
};

export default Header;

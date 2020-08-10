import React from "react";
import Typewriter from "typewriter-effect";
import classes from "./Header.module.scss";
//todo on click contact info should copy date to clipboard
import pic from "../../../assets/zoomi.jpeg";
import Social from "../../social/Social";
const Header = () => {
	return (
		<div className={`sub-container ${classes.sectionHeader}`}>
			<a className={classes.downloadResume} href="/">
				Resume <span>&darr;</span>
			</a>
			<div className={classes.header}>
				<div className={classes.imgHolder}>
					<img src={pic} alt="zamurd ali" />
				</div>
				<div className={classes.infoHolder}>
					<div className={classes.infoPersonal}>
						<p>
							Hi,
							<br /> I Wellcome you on my Portfolio
							<br /> My name is Zamurd Ali
						</p>
						<span>
							I am &nbsp;
							<Typewriter
								options={{
									strings: [
										"Learner ..!",
										"Web Developer ..!",
										"FP Gamer ..!",
										"Programmer ..!",
									],
									changeDeleteSpeed: 0.5,
									autoStart: true,
									loop: true,
								}}
							/>
						</span>
					</div>
					<div className={classes.contact}>
						<div className={classes.contact_item}>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
							>
								<path d="M17.5 2c.276 0 .5.224.5.5v19c0 .276-.224.5-.5.5h-11c-.276 0-.5-.224-.5-.5v-19c0-.276.224-.5.5-.5h11zm2.5 0c0-1.104-.896-2-2-2h-12c-1.104 0-2 .896-2 2v20c0 1.104.896 2 2 2h12c1.104 0 2-.896 2-2v-20zm-9.5 1h3c.276 0 .5.224.5.501 0 .275-.224.499-.5.499h-3c-.275 0-.5-.224-.5-.499 0-.277.225-.501.5-.501zm1.5 18c-.553 0-1-.448-1-1s.447-1 1-1c.552 0 .999.448.999 1s-.447 1-.999 1zm5-3h-10v-13h10v13z" />
							</svg>
							<span>+971-551299472</span>
						</div>
						<div className={classes.contact_item}>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
							>
								<path d="M0 3v18h24v-18h-24zm21.518 2l-9.518 7.713-9.518-7.713h19.036zm-19.518 14v-11.817l10 8.104 10-8.104v11.817h-20z" />
							</svg>
							<span>zoomi.raja@gmail.com</span>
						</div>
					</div>
					<Social />
				</div>
			</div>
		</div>
	);
};
export default Header;

import React from "react";
import classes from "./Interests.module.scss";
const Interests = () => {
	let skills = [
		"Vanilla JS",
		"Node",
		"React",
		"Redux",
		"GoLang",
		"PHP",
		"Laravel",
		"MySql",
		"Mongo",
		"Html",
		"CSS",
		"SASS",
		"Webpack",
		"OOP",
		"Git",
		"Docker",
		"API",
		"JSON",
		"Ecommerce",
		"Socket.io",
		"webRTC",
		"Microservices",
		"Video Games",
	];
	let html;
	html = skills.map((skill, i) => {
		return (
			<span key={i} className={classes.skill}>
				{skill}
			</span>
		);
	});
	return (
		<div className="sub-container">
			<h1 className="text-ac mb-2">
				<span className="text-ul">Intrests & Skills</span>
			</h1>
			<div className={classes.skills}>{html}</div>
		</div>
	);
};

export default Interests;

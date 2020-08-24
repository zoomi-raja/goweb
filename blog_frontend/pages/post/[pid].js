import axios from "axios";
import Head from "next/head";
import classes from "./Blog.module.scss";
import Layout from "../../portfolio/hoc/layout/Layout";
function markupToHtml({ body }) {
	body = body.replace(/\n/g, "<br />");

	body = body.replace(/(>|^)#{1}\s(.*?)(?=<|$)/gm, "$1<h1>$2</h1>");
	body = body.replace(/(>|^)#{2}\s(.*?)(?=<|$)/gm, "$1<h2>$2</h2>");
	//for heading h3
	body = body.replace(/(>|^)#{3}\s(.*?)(?=<|$)/gm, "$1<h3>$2</h3>");
	//to make text bold
	body = body.replace(/\*{2}(.*?)\*{2}/gm, "<strong>$1</strong>");
	//parse block of code
	body = body.replace(
		/(>|^)\`{3}(.+?)(\`{3})/gm,
		"$1<pre><code>$2</code></pre>"
	);

	//parse single line fo code
	body = body.replace(
		/(>|^|\s)\`{1}(.+?)(\`{1})/gm,
		"$1<p><code>$2</code></p>"
	);
	//conver markup for hyperlinks

	body = body.replace(
		/\[([\/\s'a-zA-Z]+?)\]\((.*?)\)/gm,
		"<a href='$2'>$1</a>"
	);
	body = body.replace(/<br \/><br \/>/g, "<br />");
	return body;
}
const Post = (props) => {
	let html = markupToHtml({ ...props.data });
	return (
		<div>
			<Head>
				<title>{props.data.title} - Post</title>
				<meta name="description" content={props.data.metatag} />
			</Head>
			<div className="container">
				<Layout header="blog">
					<section className="section">
						<h1 className="text-ac">
							<span className="text-ul">{props.data.title}</span>
						</h1>
						<div
							className={`sub-container paragraph ${classes.blog}`}
							dangerouslySetInnerHTML={{ __html: html }}
						></div>
					</section>
				</Layout>
			</div>
		</div>
	);
};
export async function getServerSideProps(ctx) {
	// Fetch data from external API
	const { pid } = ctx.params;
	// const post = await axios.get(`http://blog_api:8000/posts/${pid}/`);
	post.data = {
		id: 11,
		userId: 0,
		claps: 0,
		intro:
			"A function's this keyword behaves a little differently in JavaScript compared to other languages. It also has some differences between strict mode and non-strict mode.  In most cases, the value of this is determined by how a function is called (runti",
		title: "JavaScript `this`",
		body:
			'My first ever writing experience. We always get help from the open source community and come to the conclusion whatever we learn we should share because there is always room for improvement. I will try to demystify one of the core concept of JavaScript **this**\n## `This Keyword`\n`This` keyword is one of the most powerful mechanisms but still one of the concepts which is mostly misunderstood. to understand this we need to understand `scope`. Scope is the set of rules that controls how references to variables are available to location where you define your function on top of it, function have one more characteristic beside their scope that characteristic you can best describe as execution context. So in plain english you can not say that `this` in a method is bound to that particular object where it\'s declared. have a look\n```let obj = {\n  roles:["backend", "frontend"],\n  print:function(){\n    console.log(this.roles);\n  }\n}\nobj.print(); // ["backend", "frontend"]\nlet hold = obj.print;\nhold(); // undefined```\nyou can clearly see that when we passed reference of print method to another variable execution context has changed. we will explore in a moment how to retain this but let\'s explore where it can be useful and why we need to be expert in this.\n```let obj = {\n  roles:["backend", "frontend"],\n  print:function(){ \n    this.roles.forEach(function(role){\n      console.log(role,this.roles)\n    } )\n  }\n}\nobj.print(); // backend undefined```\nhere you will notice that inside the callback function of map `this` is no longer bound to our object. Let\'s try to set up rules how `this` works and how we can keep `this` to respond in the way as we want.\n- when inside the object method and called using that object this will remain that particular object.\n- but when the method of that object is passed to another variable (as in JavaScript other than primitive values, value is passed by reference) its context is changed so it will behave according to the new context.\n- `this` inside the function points to global context if that function is not an arrow function. in case of arrow function this will behave like lexical variable(will see shortly).\n- object created using New keyword it belongs to newly created object.\n- bind, call and apply we can force to utilize this(not discussed in this blog)\nnow lets see how we can use `this` by keeping above mentioned rules in mind to make our code work as we want\n```let obj = {\n  roles:["backend","frontend"], \n  print:function(){\n    let that = this;\n    return function(){ console.log(that.roles) };\n  }\n}\nlet hold = obj.print();\nhold(); // ["backend", "frontend"]```\nWe just used a lexical scope here to preserve this. So now it\'s easy to prove our point 3 that in the arrow function `this` behaves like a lexical variable. have a look\n```let obj = {\n  roles:["backend","frontend"],\n  print:function(){\n    return ()=>{ console.log(this.roles) }; \n  }\n}\nlet hold = obj.print();\nhold(); // ["backend", "frontend"]```\nit will work like a charm. and using same lexical this we can fix the issue in our second example\n```let obj = {\n  roles:["backend", "frontend"],\n  print:function(){\n    this.roles.forEach((role)=>{\n      console.log(role,this.roles)\n    } )\n  }\n}\nobj.print() // backend > ["backend", "frontend"]```\nnow by just keeping these rules in mind you can go to any level of complexity\n```let obj = {\n  roles:["backend", "frontend"],\n  print:function(){\n    return ()=>{\n      this.roles.forEach((role)=>{\n        console.log(role,this.roles)\n      })\n    }\n  }\n}\nlet hold = obj.print()\nhold() // backend > ["backend", "frontend"]```\nHope you enjoyed it. I focused only on this but there are relevant terminologies one need to understand to fully grasp idea. understanding lexical scope and closure and on top of it how javascript program is parsed is very important. I would recommend following two books to read.\n- [You dont know JS](https://github.com/getify/You-Dont-Know-JS)\n- [Murach’s JavaScript](https://www.murach.com/shop/murach-s-javascript-and-jquery-3rd-edition-detail)',
		createdAt: "22-Aug-2020 1:24 PM",
		updatedAt: "22-Aug-2020 1:24 PM",
	};
	// Pass data to the page via props
	return { props: post.data };
}
export default Post;

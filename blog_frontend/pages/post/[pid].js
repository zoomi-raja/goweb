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
	const post = await axios.get(`http://blog_api:8000/posts/${pid}/`);
	// Pass data to the page via props
	return { props: post.data };
}
export default Post;

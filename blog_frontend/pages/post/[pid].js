import Head from "next/head";
import Layout from "../../portfolio/hoc/layout/Layout";
import axios from "axios";
const Post = (props) => {
	console.log(props.data.body);
	props.data.body = props.data.body.replace(/\n/g, "<br />");
	props.data.body = props.data.body.replace(
		/(>|^)#{1}\s(.*?)(?=<|$)/gm,
		"$1<h1>$2</h1>"
	);
	props.data.body = props.data.body.replace(
		/(>|^)#{2}\s(.*?)(?=<|$)/gm,
		"$1<h2>$2</h2>"
	);
	//for heading h3
	props.data.body = props.data.body.replace(
		/(>|^)#{3}\s(.*?)(?=<|$)/gm,
		"$1<h3>$2</h3>"
	);
	//to make text bold
	props.data.body = props.data.body.replace(
		/\*{2}(.*?)\*{2}/gm,
		"<strong>$1</strong>"
	);
	//parse block of code
	props.data.body = props.data.body.replace(
		/(>|^)\`{3}(.+?)(\`{3})/gm,
		"$1<pre><code>$2</code></pre>"
	);
	//parse single line fo code
	props.data.body = props.data.body.replace(
		/(>|^)\`{1}(.+?)(\`{1})/gm,
		"$1<p><code>$2</code></p>"
	);
	//conver markup for hyperlinks

	props.data.body = props.data.body.replace(
		/\[([\/a-zA-Z]+?)\]\((.*?)\)/gm,
		"<a href='$2'>$1</a>"
	);
	return (
		<div>
			<Head>
				<title>{props.data.title} - Post</title>
				<meta name="description" content={props.data.metatag} />
			</Head>
			<div className="container">
				<Layout header="blog">
					<section className="section">
						<div
							className="sub-container"
							dangerouslySetInnerHTML={{ __html: props.data.body }}
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

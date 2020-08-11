import Head from "next/head";
const Post = (props) => {
	return (
		<div>
			<Head>
				<title>{props.title} - Post</title>
				<meta name="description" content={props.metatag} />
			</Head>
		</div>
	);
};
export async function getServerSideProps() {
	// Fetch data from external API
	// const res = await axi
	// const data = await res.json()

	// Pass data to the page via props
	return { props: { title: "hello" } };
}
export default Post;

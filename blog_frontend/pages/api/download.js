// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
const fs = require("fs");

export default (req, res) => {
	res.statusCode = 200;
	let message = "Not yet configured with production env";
	let filename = "resume_zamurd.pdf";
	const filepath = `${filename}`;
	if (fs.existsSync(filepath)) {
		var file = fs.readFileSync(filepath);
		res.setHeader("Content-disposition", `attachment; filename=${filename}`);
		return res.send(file);
	} else {
		res.json({ message });
	}
};

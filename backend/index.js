const express = require("express");
const app = express();
require("dotenv").config();
const cors = require("cors");

app.use(cors({ origin: "*" }));
app.use(express.json());
const morgan = require("morgan");
app.use(morgan("dev"));

app.get("/price", async (req, res) => {
  const symbols = req.query.symbols || "bitcoin,ethereum";
  const currencies = req.query.currencies || "usd";

  try {
    const url = `https://api.coingecko.com/api/v3/simple/price?ids=${symbols}&vs_currencies=${currencies}`;
    const response = await fetch(url);

    if (!response.ok) {
      return res
        .status(500)
        .json({ error: "Failed to fetch data from CoinGecko API" });
    }

    const data = await response.json();
    return res.json(data);
  } catch (error) {
    console.error(error);
    res.status(500).json({ error: "Internal Server Error" });
  }
});

app.listen(process.env.PORT || 3000, () => {
  console.log(`Server is running on port ${process.env.PORT || 3000}`);
});

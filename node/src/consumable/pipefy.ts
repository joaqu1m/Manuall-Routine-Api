import axios from "axios";

const pipeIds = {
  1: process.env.PIPEFY_CONTRATANTE_PIPE_ID,
  2: process.env.PIPEFY_PRESTADOR_PIPE_ID,
};

const getPipefyProspects = async (userType: 1 | 2) => {
  try {
    const data = axios.post(
      "https://api.pipefy.com/graphql",
      {
        query: /* GraphQL */ `{
								allCards(pipeId: ${pipeIds[userType]}) {
										edges {
												node {
														id
														fields {
																name
																value
														}
														current_phase {
																name
														}
												}
										}
								}
						}`,
      },
      { headers: { Authorization: "Bearer " + process.env.PIPEFY_API_KEY } }
    );
    return data;
  } catch (err) {
    console.log(err);
  }
};

export default getPipefyProspects;

export const data = {
  status: {
    opts: {
      Lead: 1,
      Oportunidade: 2,
      "Cliente Adquirido": 3,
      "Cliente Perdido": 4,
    },
  },
  // FORMULÁRIO INICIAL - CONTRATANTE E PRESTADOR
  "A partir de qual canal você chegou aqui?": {
    dbCol: "optCanal",
    opts: {
      "Redes Sociais": 1,
      Pesquisa: 2,
      Indicação: 3,
      "Próprio site": 4,
    },
  },
  "Nome Completo": {
    dbCol: "nome",
  },
  Email: {
    dbCol: "email",
  },
  Telefone: {
    dbCol: "fone",
    format: (value: string) => value.substring(4).replace(/[^0-9]+/g, ""),
  },
  "Você mora em:": {
    dbCol: "optCidade",
    opts: {
      "São Paulo": 1,
      "São Bernardo do Campo": 2,
      "São Caetano do Sul": 3,
      "Santo André": 4,
      Osasco: 5,
      Bauru: 6,
      Outro: 7,
    },
  },
  "Já conhece a Manuall?": {
    dbCol: "blnConheceManuall",
    opts: {
      Sim: true,
      Não: false,
    },
  },
  // LEAD CONTRATANTE
  "Qual(is) dessas áreas de serviços você está buscando?": {
    dbCol: "areaId",
    opts: {
      Jardineiro: 1,
      Pintor: 2,
      Eletricista: 3,
      Encanador: 4,
      Marceneiro: 5,
      Montador: 6,
      Gesseiro: 7,
      Nenhuma: null,
    },
  },
  "Você teria interesse em aprender algum dos serviços citados anteriormente?":
    {
      dbCol: "blnAprender",
      opts: {
        "Sim, possuo interesse.": true,
        "Não, quero apenas contratar o prestador de serviço.": false,
      },
    },
  "Você tem interesse pela Manuall?": {
    dbCol: "blnInteresseManuall",
    opts: {
      Sim: true,
      Não: false,
    },
  },
  // LEAD - PRESTADOR
  "Selecione a sua área de serviço de interesse:": {
    dbCol: "areaId",
    opts: {
      Jardineiro: 1,
      Pintor: 2,
      Eletricista: 3,
      Encanador: 4,
      Marceneiro: 5,
      Montador: 6,
      Gesseiro: 7,
      Nenhuma: null,
    },
  },
  "Você teria interesse em ensinar um pouco sobre a sua área ao outro?": {
    dbCol: "blnInteresseEnsinar",
    opts: {
      "Sim, possuo interesse.": true,
      "Não, quero apenas prestar meu serviço.": false,
    },
  },
  // OPORTUNIDADE CONTRATANTE
  "Utilizou o cupom e se tornou um cliente Contratante da Manuall?": {
    dbCol: "blnCupom",
    opts: {
      Sim: true,
      Não: false,
    },
  },
  // OPORTUNIDADE PRESTADOR
  "Utilizou o cupom e se tornou um cliente Prestador de Serviço da Manuall?": {
    dbCol: "blnCupom",
    opts: {
      Sim: true,
      Não: false,
    },
  },
  // CLIENTE PERDIDO - CONTRATANTE E PRESTADOR
  "Explique sua falta de interesse pela Manuall": {
    dbCol: "msgDesistencia",
  },
};

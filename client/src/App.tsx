import Header from "@/components/header";
import Footer from "@/components/footer";

function App() {
  return (
    <div className="w-11/12 h-full flex flex-col justify-between grow items-center">
      <Header />
      <div className="text-primary w-xl text-center flex flex-col gap-4">
        <h1 className="text-2xl font-semibold">Welcome to Ágora</h1>
        <p>
          Ágora é um fórum acadêmico onde estudantes e pesquisadores podem tirar
          dúvidas, compartilhar pesquisas e discutir ideias de forma organizada.
        </p>
      </div>
      <Footer />
    </div>
  );
}

export default App;

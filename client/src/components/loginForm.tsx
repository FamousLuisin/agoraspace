import { AlertCircle, Loader2Icon, UserPlus } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Form } from "@/components/ui/form";
import FormInput from "./formInput";
import { Button } from "./ui/button";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import z from "zod";
import { useState } from "react";
import { useAuth } from "@/context/auth-provider";

const formSchema = z.object({
  email: z.string().email({ message: "Email not is valid" }),
  password: z.string().min(8, {
    message: "Password must be at least 8 characters.",
  }),
});

interface responseSignIn {
  token: string;
}

export default function LoginForm() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const [err, setErr] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const navigate = useNavigate();
  const { setAuth } = useAuth();

  async function onSubmit(values: z.infer<typeof formSchema>) {
    setIsLoading(true);

    const url = `${import.meta.env.VITE_API_URL}/auth/signin`;

    const bodyRequest: RequestInit = {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: values.email,
        password: values.password,
      }),
    };

    try {
      const response = await fetch(url, bodyRequest);
      if (response.ok) {
        const token: responseSignIn = await response.json();
        setAuth(token.token);
        navigate("/");
      } else if (response.status === 400) {
        setErr("Failed to login. Please try again.");
      }
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
    } catch (error) {
      setErr("Failed to login. Please try again.");
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="space-y-4 w-96 border-2 rounded-2xl flex flex-col"
      >
        <div className="w-full border-b-2 p-3 flex items-center gap-2">
          <UserPlus />
          <h1 className="text-2xl font-semibold">Login</h1>
        </div>
        {err && (
          <div className="p-3">
            <div className="bg-red-900/70 border border-red-700 text-white p-4 rounded-lg">
              <div className="flex items-center gap-2">
                <AlertCircle className="h-4 w-4" />
                <p className="font-semibold">Login Error</p>
              </div>
              <p className="mt-1 text-sm">{err}</p>
            </div>
          </div>
        )}
        <FormInput label="Email" name="email" form={form} type="email" />
        <FormInput
          label="Password"
          name="password"
          form={form}
          type="password"
        />
        <div className="w-full border-b-2 p-3 flex justify-center">
          <Button className="w-full" type="submit">
            {isLoading && <Loader2Icon className="animate-spin" />}
            Submit
          </Button>
        </div>
      </form>
    </Form>
  );
}

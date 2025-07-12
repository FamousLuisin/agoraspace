import { zodResolver } from "@hookform/resolvers/zod";
import { AlertCircle, Loader2Icon, UserPlus } from "lucide-react";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import z from "zod";
import { Button } from "./ui/button";
import FormInput from "./formInput";
import { Form } from "./ui/form";
import { useAuth } from "@/context/auth-provider";

const formSchema = z
  .object({
    name: z.string().min(2, {
      message: "Name must be at least 2 characters.",
    }),
    birthday: z.string().date(),
    username: z.string().min(3, {
      message: "Username must be at least 3 characters.",
    }),
    displayname: z.string().min(3, {
      message: "Displayname must be at least 3 characters.",
    }),
    email: z.string().email({ message: "Email not is valid" }),
    password: z.string().min(8, {
      message: "Password must be at least 8 characters.",
    }),
    confirmPassword: z.string().min(8, {
      message: "Password must be at least 8 characters.",
    }),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Password don't match",
    path: ["confirmPassword"],
  })
  .refine(
    (data) => {
      const birthday = new Date(data.birthday);
      const current = new Date();
      if (current.getFullYear() - birthday.getFullYear() < 15) {
        console.log(current.getFullYear() - birthday.getFullYear() < 15);
        return false;
      }

      if (birthday.getFullYear() < 1900) {
        return false;
      }

      return true;
    },
    {
      message: "invalid date",
      path: ["birthday"],
    }
  );

interface responseSignUp {
  token: string;
}

export default function RegisterForm() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      username: "",
      displayname: "",
      email: "",
      password: "",
      confirmPassword: "",
      birthday: "",
    },
  });

  const [err, setErr] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const navigate = useNavigate();
  const { setAuth } = useAuth();

  async function onSubmit(values: z.infer<typeof formSchema>) {
    setIsLoading(true);

    const url = `${import.meta.env.VITE_API_URL}/auth/signup`;

    const bodyRequest = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: values.name,
        username: values.username,
        email: values.email,
        displayname: values.displayname,
        birthday: values.birthday,
        password: values.password,
        confirmPassword: values.confirmPassword,
      }),
    };

    try {
      const response = await fetch(url, bodyRequest);
      if (response.status === 201 || response.status === 200) {
        const token: responseSignUp = await response.json();
        setAuth(token.token);
        navigate("/");
      } else if (response.status === 400) {
        setErr("Failed to register. Please try again.");
        const err: responseSignUp = await response.json();
        console.log(err);
      }
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
    } catch (error) {
      setErr("Failed to register. Please try again.");
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="space-y-3 w-96 border-2 rounded-2xl flex flex-col"
      >
        <div className="w-full border-b-2 p-3 flex items-center gap-2">
          <UserPlus />
          <h1 className="text-2xl font-semibold">Register</h1>
        </div>
        {err && (
          <div className="p-3">
            <div className="bg-red-900/70 border border-red-700 text-white p-4 rounded-lg">
              <div className="flex items-center gap-2">
                <AlertCircle className="h-4 w-4" />
                <p className="font-semibold">Registration Error</p>
              </div>
              <p className="mt-1 text-sm">{err}</p>
            </div>
          </div>
        )}
        <FormInput label="Name" name="name" form={form} type="text" />
        <FormInput label="Username" name="username" form={form} type="text" />
        <FormInput
          label="Display name"
          name="displayname"
          form={form}
          type="text"
        />
        <FormInput label="Birthday" name="birthday" form={form} type="date" />
        <FormInput label="Email" name="email" form={form} type="email" />
        <FormInput
          label="Password"
          name="password"
          form={form}
          type="password"
        />
        <FormInput
          label="Confirm Password"
          name="confirmPassword"
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
